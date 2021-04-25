async function addPayment() {
    date = String(document.getElementById("date").value)
    _class = String(document.getElementById("class").value)
    payment = String(document.getElementById("payment").value)
    remark = String(document.getElementById("remark").value)
    if (date == "" || _class == "" || payment == "" || remark == "") {
        alert("尚未填寫完整")
        return
    }
    // payment remark class
    data = {
            "date": date,
            "class": _class,
            "payment": payment,
            "remark": remark
    }
    const res = await fetch((hostUrl + 'api/addPayment'), {
        method: 'POST',
        body: data,
        headers: {
            'Content-Type':'application/json'
        }
    })
    if (res.status == 200) {
        alert("新增成功") 
        window.location.href = hostUrl
        return
    }
    else {
        alert("新增失敗") 
        window.location.href = hostUrl + "payment"            
    }    
}