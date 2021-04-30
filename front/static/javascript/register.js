const actRegister = async (data) => {
    const res = await fetch((hostUrl + 'api/register'), {
        method: 'POST',
        body: data,
        headers: {
            'Content-Type':'application/json'
        }
    })
    if (res.status == 200) {
        alert("申請成功!") 
        window.location.href = hostUrl + "login"
        return
    }
    else {
        alert("申請失敗") 
        window.location.href = hostUrl + "register"            
    }
    return
}
function register() {
    id = String(document.getElementById("id").value)
    psd = String(document.getElementById("password").value)
    cpsd = String(document.getElementById("checkPassword").value)
    _name = String(document.getElementById("name").value)
    email = String(document.getElementById("email").value)

    if (psd != cpsd) {
        alert("密碼與確認密碼不符")  
        return          
    }
    if (id == "" || psd== "" || _name == "" || email == "") {
        alert("填寫不完全")
        return
    }
    if (validateEmail(email)){

    } else {
        alert("email格式不正確")
        return
    }
    ret = {
            "id": id,
            "password": psd,
            "name": _name,
            "email": email
    }
    actRegister(JSON.stringify(ret))
}

function validateEmail(email) {
    const re = /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(email);
  }