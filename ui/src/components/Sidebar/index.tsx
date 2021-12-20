import './sidebar.css';

const Sidebar = () => {
  return (
    <div
      className="d-flex flex-column flex-shrink-0 p-3 text-white bg-dark fixed"
      style={{ width: '280px', minHeight: '100vh' }}
    >
      <a
        href="/dashboard"
        className="d-flex align-items-center mb-3 mb-md-0 me-md-auto text-white text-decoration-none"
      >
        {/* TODO: Add SSH Management ICON */}
        <span className="fs-4">SSH Management</span>
      </a>
      <hr />
      <ul className="nav nav-pills flex-column mb-auto">
        <li className="nav-item">
          <a href="#" className="nav-link active" aria-current="page">
            {/* TODO: Add Icon */}
            Home
          </a>
        </li>
      </ul>
      <hr />
    </div>
  );
};

export default Sidebar;
