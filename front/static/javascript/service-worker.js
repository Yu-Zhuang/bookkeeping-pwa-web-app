const CACHE_NAME = "gokeep_cache_v1"
// 緩存內容
self.addEventListener('install', async event => {
    console.log('install', event)
    // 開啟一個緩存
    const cache = await caches.open(CACHE_NAME)
    // 儲存資料在cache中
    await cache.addAll([
        '/',
        'static/javascript/manifest.json'
    ])
    await self.skipWaiting()
})

// 清除舊緩存
self.addEventListener('activate', async event => {
    console.log('activate', event)
    // 獲取所有cache的keys
    const keys = await caches.keys()
    // 走訪各個cache
    keys.forEach(k => {
        // 如果是舊的key則刪除
        if(k !== CACHE_NAME){
            caches.delete(k)
        }
    })
    await self.clients.claim()
})

self.addEventListener('fetch', event => {
    const req = event.request
    event.respondWith(netWorkFirst(req))
})

async function netWorkFirst(req) {
    // 抓出請求
    try{
        const fresh = await fetch(req)
        return fresh
    } catch(e) {
        // 失敗了: 去讀取暫存
        console.log(`讀緩存`)
        const cache = await caches.open(CACHE_NAME) // 先打開瀏覽器中的暫存
        const cached = await caches.match(req) // 找出路徑對應的cache並回傳該暫存
        return cached
    }
    
}