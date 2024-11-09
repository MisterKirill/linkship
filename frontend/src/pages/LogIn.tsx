import { FormEvent, useState } from 'react'

function LogIn() {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')

  const logIn = (e: FormEvent) => {
    e.preventDefault()

    console.log(username, password)
  }

  return (
    <>
      <h1>Welcome back!</h1>

      <form className="form" onSubmit={logIn}>
        <input type="text" placeholder="Username" className="text-field" onChange={(e) => setUsername(e.target.value)} required />
        <input type="password" placeholder="Password" className="text-field" onChange={(e) => setPassword(e.target.value)} required />
        <input type="submit" value="Log In" className="button" />
      </form>
    </>
  )
}

export default LogIn
