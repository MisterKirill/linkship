import { Link } from 'react-router-dom'

function NotFoundPage() {
  return (
    <>
      <h1>404</h1>

      <span>
        The requested URL was not found.<br />
        You can go <Link to="/" className="text-link">home</Link> or <a href="https://github.com/MisterKirill/linkship/issues" className="text-link">create an issue</a> on GitHub if you think there is a mistake.
      </span>
    </>
  )
}

export default NotFoundPage
