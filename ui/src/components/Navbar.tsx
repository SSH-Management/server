import { FC } from 'react';
import Avatar from 'react-avatar';
import { Container, Nav, NavDropdown, Navbar as BsNavbar } from 'react-bootstrap';

const Navbar: FC = () => {
  return (
    <BsNavbar bg="dark" expand="lg" variant="dark" sticky="top">
      <Container>
        <BsNavbar.Toggle aria-controls="navbarScroll" />
        <BsNavbar.Collapse id="navbarScroll">
          <Nav className="me-auto my-2 my-lg-0" style={{ maxHeight: '100px' }} navbarScroll />
          <Nav className="my-2 my-lg-0 me-2" style={{ maxHeight: '100px' }} navbarScroll>
            <NavDropdown
              id="nav-dropdown-dark-example"
              title={<Avatar name="Foo Bar" size={'30'} round="20px" />}
              menuVariant="dark"
            >
              <NavDropdown.Item href="#action/3.1">Dashboard</NavDropdown.Item>
              <NavDropdown.Item href="#action/3.2">Profile</NavDropdown.Item>
              <NavDropdown.Divider />
              <NavDropdown.Item href="#action/3.4">Logout</NavDropdown.Item>
            </NavDropdown>
          </Nav>
        </BsNavbar.Collapse>
      </Container>
    </BsNavbar>
  );
};

export default Navbar;
