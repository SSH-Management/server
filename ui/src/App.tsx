import { FC, lazy, Suspense } from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';

const Layout = lazy(() => import('./components/Layout'));
const Login = lazy(() => import('./Pages/Login'));

const App: FC = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route
          path="/"
          element={
            <Suspense fallback={<>...</>}>
              <Layout />
            </Suspense>
          }
        ></Route>
        <Route
          path="/login"
          element={
            <Suspense fallback={<>...</>}>
              <Login />
            </Suspense>
          }
        />
      </Routes>
    </BrowserRouter>
    // <div className="d-flex">
    //   <Sidebar />
    //   <div className="flex-grow-1">
    //     <Navbar />
    //     {/* <Login /> */}
    //     {/* <Container> */}
    //     {/* <Row>
    //         <DataTable
    //           direction={Direction.LTR}
    //           expandOnRowClicked
    //           fixedHeader
    //           fixedHeaderScrollHeight="500px"
    //           highlightOnHover
    //           pagination
    //           persistTableHead
    //           responsive
    //           selectableRows
    //           selectableRowsHighlight
    //           columns={[
    //             {
    //               name: 'Name',
    //             },
    //             {
    //               name: 'Surname',
    //             },
    //             {
    //               name: 'Email',
    //             },
    //           ]}
    //           data={[]} // subHeaderAlign={Alignment.RIGHT}
    //           // subHeaderWrap
    //         />
    //       </Row>
    //     </Container> */}
    //   </div>
    // </div>
  );
};

export default App;
