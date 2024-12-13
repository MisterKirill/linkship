import { FormEvent, useEffect, useState } from "react"
import { Link, useNavigate } from "react-router-dom"
import { authenticate } from "../../utils/users"
import classes from "./style.module.css"

function Login() {
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  const navigate = useNavigate()

  const onSubmit = async (event: FormEvent) => {
    event.preventDefault()

    if (await authenticate({ username, password })) {
      navigate("/")
    }
  }

  useEffect(() => {
    document.title = "Log In - Linkship"
  }, [])
  
  return (
    <div className={classes.wrapper}>
      <div>
        <h1 className={classes.heading}>Welcome to the <span className="text-brand">Linkship</span>!</h1>
        <span>Don't have an account yet? <Link to="/register" className="text-link">Register</Link>.</span>
      </div>

      <form className="form" onSubmit={onSubmit}>
        <input
          type="text"
          placeholder="Username"
          className="input"
          onChange={(e) => setUsername(e.target.value)}
          required />

        <input
          type="password"
          placeholder="Password"
          className="input"
          onChange={(e) => setPassword(e.target.value)}
          required />

        <button type="submit" className="btn">
          Log In
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="currentColor" viewBox="0 0 16 16">
            <path fillRule="evenodd" d="M4.646 1.646a.5.5 0 0 1 .708 0l6 6a.5.5 0 0 1 0 .708l-6 6a.5.5 0 0 1-.708-.708L10.293 8 4.646 2.354a.5.5 0 0 1 0-.708"/>
          </svg>
        </button>
      </form>
    </div>
  )
}

export default Login
