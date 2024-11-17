import { FormEvent, useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { authenticate } from '../utils/users'

function LoginPage() {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const navigate = useNavigate()

  const onSubmit = async (e: FormEvent) => {
    e.preventDefault()

    if (await authenticate(username, password)) {
      navigate('/')
    }
  }

  return (
    <>
      <h1>Welcome back!</h1>

      <form className="form" onSubmit={onSubmit}>
        <input type="text" placeholder="Username" className="text-field" onChange={(e) => setUsername(e.target.value)} required />
        <input type="password" placeholder="Password" className="text-field" onChange={(e) => setPassword(e.target.value)} required />
        <input type="submit" value="Log In" className="button" />
      </form>

      <p>Don't have an account yet? <Link to="/register" className="text-link">Register</Link>.</p>
    </>
  )
}

export default LoginPage
