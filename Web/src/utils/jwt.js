export function parseToken(token){
    let payLoad = token.split(".")[1]
    let adminInfo = JSON.parse(decodeURIComponent(window.atob(payLoad.replace(/-/g, "+").replace(/_/g, "/"))))
    return adminInfo
}