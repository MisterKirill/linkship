import { Link } from './links'

export interface User {
  username: string
  display_name: string
  bio: string
  links: Link[]
}

export async function authenticate(
  username: string,
  password: string,
  isRegister: boolean = false
): Promise<boolean> {
  const res = await fetch(
    isRegister ? `${import.meta.env.VITE_BACKEND_URL}/users` : `${import.meta.env.VITE_BACKEND_URL}/users/login`,
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

export async function getUser(username: string): Promise<User | null> {
  const res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/users/${username}`)

  const json = await res.json()

  if (res.status == 200) {
    return json
  } else {
    return null
  }
}
