import { FC } from 'react';
import { Outlet } from 'react-router-dom';
import Navbar from './Navbar';
import Sidebar from './Sidebar';

const Layout: FC = () => {
  return (
    <div className="d-flex">
      <Sidebar />
      <div className="flex-grow-1">
        <Navbar />
        <Outlet />
        {/* <Login /> */}
        {/* <Container> */}
        {/* <Row>
            <DataTable
              direction={Direction.LTR}
              expandOnRowClicked
              fixedHeader
              fixedHeaderScrollHeight="500px"
              highlightOnHover
              pagination
              persistTableHead
              responsive
              selectableRows
              selectableRowsHighlight
              columns={[
                {
                  name: 'Name',
                },
                {
                  name: 'Surname',
                },
                {
                  name: 'Email',
                },
              ]}
              data={[]} // subHeaderAlign={Alignment.RIGHT}
              // subHeaderWrap
            />
          </Row>
        </Container> */}
      </div>
    </div>
  );
};

export default Layout;
