import { Link } from 'react-router-dom'

function NotFound() {
  return (
    <>
      <h1>404</h1>
      
      <span>This page was not found.</span>
      
      <span>
        You can go <Link to="/" className="link">home</Link> or <a href="https://github.com/MisterKirill/linkship/issues" className="link">create an issue</a> on GitHub if you think that something went wrong.
      </span>
    </>
  )
}

export default NotFound
