import { Switch, Route } from 'react-router';
import {
  CartList,
  EmailRegister,
  Home,
  Login,
  NotFound,
  OrderConfirm,
  OrderHistory,
  PasswordReset,
  ProductDetail,
  Register,
  UserMyPage,
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
      <Route exact path={'/product/:productId?'} component={ProductDetail} />
      <Auth>
        <Switch>
          <Route exact path={'/user/mypage'} component={UserMyPage} />
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
