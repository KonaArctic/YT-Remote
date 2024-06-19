package main
import "io"
import "net/http"
import "net/url"
import "os"

func main( ) {
    var err error
    var stream map[ string ]http.ResponseWriter = map[ string ]http.ResponseWriter{ }
    var server http.Server = http.Server{
        Addr : "[::]:80" ,
        Handler : http.HandlerFunc( func( respond http.ResponseWriter , request * http.Request ){
            respond.Header( )[ "Access-Control-Allow-Origin" ] = [ ]string{ "*" }
            var qryurl * url.URL
            qryurl , err = url.Parse( request.URL.RawQuery )
            if  err != nil ||
                qryurl.Scheme == "" {
                switch request.Method {
                    case http.MethodHead :
                        respond.WriteHeader( http.StatusOK )
                    case http.MethodGet :
                        respond.WriteHeader( http.StatusOK )
                        respond.( http.Flusher ).Flush( )
                        stream[ request.URL.RawQuery ] = respond
                        <- make( chan bool , 1 )    // FIXME Leaky
                    case http.MethodPost :
                        if _ , ok := stream[ request.URL.RawQuery ] ; ok {
                            _ , err = io.Copy( stream[ request.URL.RawQuery ] , request.Body )
                            if err == nil {
                                stream[ request.URL.RawQuery ].( http.Flusher ).Flush( )
                                respond.WriteHeader( http.StatusOK )
                                return
                            }
                        }
                        respond.WriteHeader( http.StatusNotFound )
                    default :
                        respond.WriteHeader( http.StatusBadRequest )
                }
            } else {
//respond.WriteHeader( http.StatusBadRequest )
//return
                var qryrsp * http.Response
                qryrsp , err = http.DefaultTransport.RoundTrip( & http.Request{
                    Method : request.Method ,
                    URL : qryurl ,
                    Header : request.Header ,
                    Body : request.Body ,
                } )
                if err != nil {
                    respond.WriteHeader( http.StatusBadGateway )
                    return }
                for key , value := range qryrsp.Header {
                    if key != "Access-Control-Allow-Origin" {
                        respond.Header( )[ key ] = value } }
                respond.WriteHeader( qryrsp.StatusCode )
                _ , _ = io.Copy( respond , qryrsp.Body )
            }
        } ) ,
    }
    if len( os.Args ) >= 2 {
        server.Addr = os.Args[ 1 ] }
    err = server.ListenAndServe( )
    os.Exit( 1 )
}
