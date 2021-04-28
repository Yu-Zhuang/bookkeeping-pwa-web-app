async function getPaymentRecord() {
    container = document.getElementById('record')
    const res = await fetch((hostUrl + 'api/getPaymentHistory'), {
        method: 'GET',
    })
    if (res.status == 200) {
        myJson = await res.json()
        data = myJson.data
        if (data.length < 1) {
            record = preproRecord('空', '空', '空', '空')
        } else {
            for (let i = data.length-1; i >= 0; i--) {
                record = preproRecord(data[i].date, data[i].class, data[i].payment, data[i].remark)
                container.innerHTML += record
            }
        }
        return
    }
    else {
        record = preproRecord('空', '空', '空', '空')
        container.innerHTML += record
    }
}
function preproRecord(date, _class, payment, remark) {
    return `
    <div class="container" style="margin-top: 10px; border-bottom: 1px dashed black;">
        <div id="date" class="ele">${date}</div>
        <div id="class" class="ele">${_class}</div> 
        <div id="payment" class="ele">${payment}</div>  
        <div id="remark" class="ele">${remark}</div>  
    </div>
    `
}