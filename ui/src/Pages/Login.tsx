import { FC } from 'react';
import './Login.css';

const Login: FC = () => {
  return (
    <div className="text-center wrapper">
      <div className="form-signin">
        <form>
          <img
            className="mb-4"
            src="../assets/brand/bootstrap-logo.svg"
            alt=""
            width="72"
            height="57"
          />
          <h1 className="h3 mb-3 fw-normal">Please sign in</h1>

          <div className="form-floating">
            <input
              type="email"
              className="form-control mb-4"
              id="floatingInput"
              placeholder="name@example.com"
            />
            <label htmlFor="floatingInput">Email address</label>
          </div>
          <div className="form-floating">
            <input
              type="password"
              className="form-control"
              id="floatingPassword"
              placeholder="Password"
            />
            <label htmlFor="floatingPassword">Password</label>
          </div>
          <button className="w-100 btn btn-lg btn-primary" type="submit">
            Sign in
          </button>
          <p className="mt-5 mb-3 text-muted">&copy; SSH Management 2021</p>
        </form>
      </div>
    </div>
  );
};

export default Login;
