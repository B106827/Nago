import { useSelector } from 'react-redux';
import { getIsLogedIn } from './reducks/users/selectors';

const Auth = ({ children }) => {
  const selector = useSelector((state) => state);
  const isLogedIn = getIsLogedIn(selector);

  if (!isLogedIn) {
    return <></>;
  } else {
    return children;
  }
};

export default Auth;
