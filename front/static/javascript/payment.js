function addPayment() {
    date = String(document.getElementById("date").value)
    _class = String(document.getElementById("class").value)
    payment = String(document.getElementById("class").value)
    remark = String(document.getElementById("class").value)
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
}