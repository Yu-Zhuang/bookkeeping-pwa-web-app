const CACHE_NAME = "gokeep_cache_v3"
// 緩存內容
self.addEventListener('install', async event => {
    console.log('install', event)
    // 開啟一個緩存
    const cache = await caches.open(CACHE_NAME)
    // 儲存資料在cache中
    await cache.addAll([
        '/error',
        'static/javascript/manifest.json',
        'static/image/apple-touch-icon.png',
        'static/image/favicon-32x32.png',
        'static/image/favicon-16x16.png',
        'static/image/favicon.ico',
        'static/css/reset.css',
        'static/css/main.css',
        'static/css/index.css',
        'static/image/inactive-chart.svg',
        'static/image/inactive-target.svg',
        'static/image/inactive-addNew.svg',
        'static/image/inactive-showRecord.svg',
        'static/image/inactive-userProfile.svg',
        'static/javascript/conf.js',
        'static/javascript/index.js',
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
        const cached = await caches.match(req+'error') // 找出路徑對應的cache並回傳該暫存
        return cached
    }
    
}