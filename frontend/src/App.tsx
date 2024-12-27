import { Outlet } from 'react-router'
import { Link } from 'react-router'
import './App.css'
import 'beercss'

function App() {

  return (
    <>
      <nav className="left drawer l">
        <header>
          <nav>
            <h6>Four rooms</h6>
          </nav>
        </header>
        <div>
          <i>search</i>
          <Link to="/search">Search</Link>
        </div>
        <div>
          <i>home</i>
          <Link to="/hotels">Hotels</Link>
        </div>
      </nav>

      <nav className="top s m left-align">
        <button data-ui="#menu" className='transparent circle'>
          <i>menu</i>
          <menu id="menu" className='no-wrap'>
            <Link to="/search">Search</Link>
            <Link to="/hotels">Hotels</Link>
          </menu>
        </button>
      </nav>

      <main className="responsive">
        <Outlet />
      </main>
    </>
  )
}

export default App
