import { Outlet } from 'react-router'
import { Link } from 'react-router'
import './App.css'
import 'beercss'

function App() {

  return (
    <>
      <nav class="left drawer l">
        <header>
          <nav>
            <h6>Four rooms</h6>
          </nav>
        </header>
        <a>
          <i>home</i>
          <Link to="/hotels">Hotels</Link>
        </a>
        {/* <a>
          <i>search</i>
          <div>Search</div>
        </a> */}
      </nav>

      <nav class="top s m left-align">
        <button data-ui="#menu" className='transparent circle'>
          <i>menu</i>
          <menu id="menu" className='no-wrap'>
            <Link to="/hotels">Hotels</Link>
            {/* <a>Search</a> */}
          </menu>
        </button>
      </nav>

      <main class="responsive">
        <Outlet />
      </main>
    </>
  )
}

export default App
