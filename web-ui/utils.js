
function getBaseURI(){
    let uri = document.baseURI;
    let schemeTerm = uri.indexOf('://') + 3;
    return uri.substring(0, uri.indexOf('/', schemeTerm) + 1);
}

