import { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { getUser, User } from '../../utils/users'
import './style.css'

function ProfilePage() {
  const [user, setUser] = useState<User | null>(null)
  const [isLoading, setIsLoading] = useState<boolean>(true)
  const [isFailed, setIsFailed] = useState<boolean>(false)
  const { username } = useParams()

  useEffect(() => {
    if (username) {
      getUser(username)
        .then(user => {
          setUser(user)

          if (user) {
            document.title = `${user.display_name || user.username} - Linkship`
          }
        })
        .catch(() => setIsFailed(true))
        .finally(() => setIsLoading(false))
    }
  }, [])

  if (isLoading) {
    return (
      <span className="loader"></span>
    )
  }

  if (isFailed) {
    return (
      <>
        <h1>Oops!</h1>

        <p>Failed to get user! Please try again in a few minutes.</p>
      </>
    )
  }

  if (!user) {
    return (
      <>
        <h1>Oops!</h1>

        <p>Looks like user <b>{username}</b> does not exist.</p>
      </>
    )
  }

  return (
    <>
      <h1 className="display-name">{user.display_name || user.username}</h1>

      {user.bio && (
        <span className="bio">
          {user.bio}
        </span>
      )}

      {user.links.length > 0 ? (
        <div className="link-buttons-list">
          {user.links.map(link => {
            return (
                <a href={link.url} key={link.id} className="link-button">
                  {link.name}
                </a>
            )
          })}
        </div>
      ) : (
        <span className="text-muted">Looks like <b>{user.username}</b> hasn't added links yet!</span>
      )}
    </>
  )
}

export default ProfilePage
