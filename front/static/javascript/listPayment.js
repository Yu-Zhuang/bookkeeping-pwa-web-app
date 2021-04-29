async function getPaymentRecord() {
    container = document.getElementById('record')
    const res = await fetch((hostUrl + 'api/getPaymentHistory'), {
        method: 'GET',
    })
    if (res.status == 200) {
        myJson = await res.json()
        data = myJson.data
        if (data == null || data.length <= 0) {
            record = preproRecord('空', '空', '空', '空', null)
            container.innerHTML += record
        } else {
            for (let i = data.length-1; i >= 0; i--) {
                record = preproRecord(data[i].date, data[i].class, data[i].payment, data[i].remark, data[i]._id)
                container.innerHTML += record
            }
        }
    }
    else {
        record = preproRecord('空', '空', '空', '空', null)
        container.innerHTML += record
    }
}

function preproRecord(date, _class, payment, remark, _id) {
    return `
    <div class="container" style="margin-top: 10px; border-bottom: 1px dashed black;">
        <div id="date" class="ele">${date}</div>
        <div id="class" class="ele">${_class}</div> 
        <div id="payment" class="ele">${payment}</div>  
        <div id="remark" class="ele">${remark}</div>
        <div class="ele"><img src="static/image/trash.svg" onclick="deletItem(${_id})"></div>
    </div>
    `
}

async function deletItem(_id) {
    let req = "api/deletRecord/" + String(_id)
    const res = await fetch((hostUrl + req), {
        method: 'GET',
    })
    if (res.status == 200) {
        alert('刪除成功')
    } else {
        alert('刪除失敗')
    }
    window.location.href = hostUrl + "list-payment"
}