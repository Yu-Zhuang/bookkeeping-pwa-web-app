const register = async () => {
    const data = await getData()
    const res = await fetch((hostUrl + 'api/register'), {
        method: 'POST',
        body: data,
        headers: {
            'Content-Type':'application/json'
        }
    })
    //const myJson = await res.json()
    if (res.status == 200) {
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
    cpsd = String(document.getElementById("checkPassword").value)
    _name = String(document.getElementById("name").value)
    email = String(document.getElementById("email").value)

    if (psd != cpsd) {
        return null, false
    }
    ret = {
            "id": id,
            "password": psd,
            "name": _name,
            "email": email
    }
    return JSON.stringify(ret), true
}