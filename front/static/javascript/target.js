async function getAverage() {
    m_avg = document.getElementById('m-avg')
    d_avg = document.getElementById('d-avg')
    const res = await fetch((hostUrl + 'api/getAverage'), {
        method: 'GET',
    })
    if (res.status == 200) {
        myJson = await res.json()
        data = myJson.data
        m_avg.innerHTML += data.monthAvg
        d_avg.innerHTML += data.dayAvg
        return
    }
    else {
        m_avg.innerHTML += 'no data'
        d_avg.innerHTML += 'no data'
    }
}