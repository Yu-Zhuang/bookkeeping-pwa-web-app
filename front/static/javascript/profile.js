async function getProfile() {
    email = document.getElementById('email')
    _name = document.getElementById('_name')
    const res = await fetch((hostUrl + 'api/getProfile'), {
        method: 'GET',
    })
    if (res.status == 200) {
        myJson = await res.json()
        email.innerHTML += myJson.email
        _name.innerHTML += myJson.name
        return
    }
    email.innerHTML += `(空)`
    _name.innerHTML += `(空)`
}
async function logout() {
    const res = await fetch((hostUrl + 'api/logOut'), {
        method: 'POST',
    })
    window.location.href = hostUrl + 'login'
    return
}