const hostUrl = "https://bookkepping.herokuapp.com/"
//const hostUrl = "http://127.0.0.1:8080/"

// pwa
window.addEventListener('load', async () =>{
    if ('serviceWorker' in navigator) {
        try {
            const regi = await navigator.serviceWorker.register("zService-worker.js")
        } catch(e) {
            console.log(`sw註冊失敗`)
        }
    }
})
// // 跳提醒授權: default(未授權), granted(已授權), denid(鎖)
// if(Notification.permission === 'default') {
//     Notification.requestPermission()
// }

if(!navigator.onLine) {
    if (window.location.href != hostUrl + "error") {
        window.location.href =  hostUrl + "error"   
    } 
}

window.addEventListener('online', ()=>{
    if (window.location.href != hostUrl) {
        window.location.href = hostUrl
    } 
})

function toHome() {
    window.location.href = hostUrl
}