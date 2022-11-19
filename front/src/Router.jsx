import { Switch, Route } from 'react-router';
import {
  CartList,
  OrderConfirm,
  OrderHistory,
  ProductDetail,
  Login,
  Register,
  EmailRegister,
  PasswordReset,
  Home,
  NotFound,
} from './templates';
import Auth from './Auth';

const Router = () => {
  return (
    <Switch>
      <Route exact path={'/'} component={Home} />
      <Route path={'/register/:tmpId?'} component={Register} />
      <Route exact path={'/email_register'} component={EmailRegister} />
      <Route exact path={'/login'} component={Login} />
      <Route exact path={'/password_reset'} component={PasswordReset} />
      <Route exact path={'/product'} component={ProductDetail} />
      <Auth>
        <Switch>
          <Route exact path={'/cart'} component={CartList} />
          <Route exact path={'/order/confirm'} component={OrderConfirm} />
          <Route exact path={'/order/history'} component={OrderHistory} />
          <Route component={NotFound} />
        </Switch>
      </Auth>
    </Switch>
  );
};
export default Router;
