import { FormEvent, useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'

function Register() {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const navigate = useNavigate()

  const register = async (e: FormEvent) => {
    e.preventDefault()

    const res = await fetch('http://localhost:8080/users', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        username,
        password
      })
    })

    const json = await res.json()
    console.log(json)

    if (res.status == 200) {
      localStorage.setItem('token', json.token)
      navigate('/')
    } else {
      alert(json.message)
    }
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
