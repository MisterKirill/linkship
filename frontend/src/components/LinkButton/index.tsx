import React from 'react'
import './style.css'

function LinkButton({ backgroundColor, href, children }: {
  backgroundColor: string,
  href: string,
  children: React.ReactNode
}) {
  return (
    <a href={href} className="link">
      <button className="link-button" style={{
        backgroundColor: '#' + backgroundColor
      }}>
        {children}
      </button>
    </a>
  )
}

export default LinkButton
