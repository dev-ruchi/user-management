import { Outlet } from "react-router-dom";

function App() {
  return (
    <>
      <div className="container mx-auto mt-10">
        <Outlet />
      </div>
    </>
  );
}

export default App;
