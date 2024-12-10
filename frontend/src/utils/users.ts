interface Credentials {
  username: string
  password: string
}

export async function authenticate(credentials: Credentials, isRegister: boolean = false): Promise<boolean> {
  const res = await fetch(import.meta.env.VITE_BACKEND_URL + (isRegister ? "/users" : "/login"), {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(credentials)
  })

  const json = await res.json()

  if (res.status != 200) {
    alert(json.message)
    return false
  }

  localStorage.setItem("token", json.token)
  localStorage.setItem("username", credentials.username)

  return true
}
