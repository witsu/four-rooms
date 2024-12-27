import { useEffect, useState } from 'react'
import { useSearchParams } from "react-router";
import { formatTomorrowDate, formatDate } from './helpers';

function SearchPage() {
  const [rooms, setRooms] = useState([]);
  const [searchParams, setSearchParams] = useSearchParams();
  const today = new Date();
  const start = searchParams.get("start") ?? formatDate(today);
  const end = searchParams.get("end") ?? formatTomorrowDate(today);
  const location = searchParams.get("location") ?? 'New York';

  function search(event) {
    event.preventDefault();
    const formData = new FormData(event.target);
    const params = {
      start: formData.get("start"),
      end: formData.get("end"),
      location: formData.get("location"),
    };
    setSearchParams(params);
  }

  useEffect(() => {
    if (!start || !end || !location) {
        return;
    }
    const fetchSearch = () => {
        fetch(`http://localhost:8080/search?start=${start}&end=${end}&location=${location}`)
          .then(response => response.json())
          .then(data => setRooms(data))
          .catch(error => console.error('Error fetching data:', error))
    }
    fetchSearch();
  }, [start, end, location]);

  return (
    <>
      <form onSubmit={search}>
        <div className="field label border">
          <input type="date" name="start" defaultValue={start} />
          <label>Start Date:</label>
        </div>
        <div className="field label border">
          <input type="date" name="end" defaultValue={end} min={start} />
          <label>End Date:</label>
        </div>
        <div className="field label border">
          <input type="text" name="location" defaultValue={location} />
          <label>Location:</label>
        </div>
        <button type="submit">Search</button>
      </form>
      
      <h3>Search results</h3>
      {rooms.length === 0 && <p>No rooms available. Try changing dates or location.</p>}
      {rooms.map(room => (
        <article key={room.id}>
            <h5>from the movie "{room.title}"</h5>
            <div>{room.type}</div>
            <div>Size: {room.size} m2</div>
            <div>{room.description}</div>
            <nav>
                <button>Book</button>
            </nav>
        </article>
      ))}
    </>
  )
}

export default SearchPage
