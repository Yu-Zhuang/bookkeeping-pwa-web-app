const getUser = async () => {
    const res = await fetch((hostUrl + 'api/getUser'), {
        method: 'GET'
    })
    if (res.status < 400) {
        return
    }
    else {
        window.location.href = hostUrl + "login"            
    }
    return
}