const login = async () => {
    const data = await getData()
    const res = await fetch((hostUrl + 'api/login'), {
        method: 'POST',
        body: data,
        headers: {
            'Content-Type':'application/json'
        }
    })
    if (res.status == 200) {
        alert("登入成功!") 
        window.location.href = hostUrl
        return
    }
    else {
        alert("錯誤的帳號密碼") 
        window.location.href = hostUrl + "login"            
    }
    return
}
function getData() {
    id = String(document.getElementById("id").value)
    psd = String(document.getElementById("password").value)
    ret = {
            "id": id,
            "password": psd
        }
    return JSON.stringify(ret)
}