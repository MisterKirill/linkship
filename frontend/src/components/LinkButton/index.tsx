import React from 'react'
import './style.css'

function LinkButton({ backgroundColor, href, children }: { backgroundColor: string, href: string, children: React.ReactNode }) {
  return (
    <a href={href} className="link">
      <button className="link-button" style={{
        backgroundColor: '#' + backgroundColor,
        boxShadow: `5px 5px 0 #${backgroundColor}40`
      }}>
        {children}
      </button>
    </a>
  )
}

export default LinkButton
