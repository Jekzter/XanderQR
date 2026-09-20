import { Routes, Route } from "react-router";
import HomePages from "@app/pages/home";

function App() {

  return (
    <Routes>
      <Route path="/" element={<HomePages />} />
    </Routes>
  )
}

export default App
