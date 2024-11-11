import { FormEvent, useState } from 'react'
import { Link } from 'react-router-dom'

function Register() {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')

  const register = (e: FormEvent) => {
    e.preventDefault()

    console.log(username, password)
  }

  return (
    <>
      <h1>Welcome to Linkship!</h1>

      <form className="form" onSubmit={register}>
        <input type="text" placeholder="Username" className="text-field" onChange={(e) => setUsername(e.target.value)} required />
        <input type="password" placeholder="Password" className="text-field" onChange={(e) => setPassword(e.target.value)} required />
        <input type="submit" value="Register" className="button" />
      </form>

      <p>Already have an account? <Link to="/login" className="text-link">Log In</Link>.</p>
    </>
  )
}

export default Register
