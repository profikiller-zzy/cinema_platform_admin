export function parseToken(token){
    let payLoad = token.split(".")[1]
    let userInfo = JSON.parse(decodeURIComponent(window.atob(payLoad.replace(/-/g, "+").replace(/_/g, "/"))))
    return userInfo
}