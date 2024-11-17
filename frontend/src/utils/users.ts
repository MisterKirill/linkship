export async function authenticate(
  username: string,
  password: string,
  isRegister: boolean = false
): Promise<boolean> {
  const res = await fetch(
    isRegister ? 'http://localhost:8080/users' : 'http://localhost:8080/users/login',
    {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        username,
        password
      })
    }
  )

  const json = await res.json()

  if (res.status == 200) {
    localStorage.setItem('token', json.token)
    localStorage.setItem('username', username)
    return true
  } else {
    alert(json.message)
    return false
  }
}
