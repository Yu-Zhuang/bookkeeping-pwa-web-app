async function getAverage() {
    m_avg = document.getElementById('m-avg')
    const res = await fetch((hostUrl + 'api/getMonthAverage'), {
        method: 'GET',
    })
    if (res.status == 200) {
        myJson = await res.json()
        data = myJson.data
        m_avg.innerHTML += data.monthAvg
    }
    else {
        m_avg.innerHTML += 'no data'
    }
    getDayAverage()
}

async function getDayAverage() {
    d_avg = document.getElementById('d-avg')
    const res = await fetch((hostUrl + 'api/getDayAverage'), {
        method: 'GET',
    })
    if (res.status == 200) {
        myJson = await res.json()
        data = myJson.data
        d_avg.innerHTML += data.dayAvg
    }
    else {
        d_avg.innerHTML += 'no data'
    }
}