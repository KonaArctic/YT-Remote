// Twitter
export class twitter extends DocumentFragment {
    static get match( ) {
        return "^(https?://)?(twitter|x)\\.com/[^/]+/status/[0-9]+"
    }
    constructor( vidlnk ) { // Broken :(
        super( );
        this.append( window.document.createElement( "blockquote" ) );
        this.children[ 0 ].className = "twitter-tweet";
        this.children[ 0 ].setAttribute( "data-media-max-width" , "560" );
        this.children[ 0 ].append( window.document.createElement( "a" ) );
        this.children[ 0 ].children[ 0 ].href = vidlnk;
        this.append( window.document.createElement( "script" ) );
        this.children[ 1 ].src = "https://platform.twitter.com/widgets.js"
    }
    // TODO Implement
    play( ) {
        return;
    }
    pause( ) {
        return;
    }
    seek( value ) {
        return;
    }
}

// Youtube
let ytatag = window.document.createElement( "script" );
ytatag.src = "https://www.youtube.com/iframe_api";
ytatag.async = true;
window.document.head.append( ytatag );
export class youtube extends DocumentFragment {
    static get match( ) {
        return "^(https?://)?(((www|m)\\.)?youtube\\.com/watch\\?v=|youtu.be/)";
    }
    constructor( vidlnk ) {
        super( );
        let player = window.document.createElement( "div" );
        this.append( player );
        this.player = new YT.Player( player , {
            height : window.innerHeight ,
            width : window.innerWidth ,
            videoId: vidlnk.match("(https?://)?(((www|m)\\.)?youtube\\.com/watch\\?v=|youtu.be/)([^&?/]+)")[5] ,
            playerVars : {
                playsinline : 1
            } ,
            events: {
                onReady : function( event ){
                    event.target.playVideo( );
                } ,
                onStateChange : null ,
            } ,
        } );
    }
    play( ) {
        this.player.playVideo( );
    }
    pause( ) {
        this.player.pauseVideo( );
    }
    seek( value ) {
        this.player.seekTo( this.player.getCurrentTime( ) + value , true );
    }
}

export default [
    twitter ,
    youtube ,
];

