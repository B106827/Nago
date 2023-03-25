import { useCallback, useMemo, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { getMyCartList } from '../reducks/users/selectors';
import { makeStyles } from '@material-ui/styles';
import { CartListItem } from '../components/Products';
import List from '@material-ui/core/List';
import Divider from '@material-ui/core/Divider';
import { PrimaryButton, TextDetail } from '../components/UIkit';
import { push } from 'connected-react-router';
import config from '../config/base';
//import { orderProduct } from '../reducks/products/operations';

const OrderConfirm = () => {
  const classes  = useStyles();
  const dispatch = useDispatch();
  const selector = useSelector((state) => state);
  const cartList = getMyCartList(selector);
  useEffect(() => {
    if (!cartList || cartList.length === 0) {
      dispatch(push('/'));
    }
  })

  const subtotal = useMemo(() => {
    return cartList.reduce((sum, cart) => {
      return sum += (cart.product.price * cart.num);
    }, 0);
  }, [cartList]);

  const shippingFee = 0;
  const tax         = subtotal * config.taxRate;
  const total       = subtotal + shippingFee + tax;

  const order = useCallback(() => {
    //   dispatch(orderProduct(cartList, total));
  }, [dispatch, cartList, total]);

  return (
    <section className={classes.topSection}>
      <h2 className={classes.topSectionTitle}>注文の確認</h2>
      <div className='p-grid__row'>
        <div className={classes.detailBox}>
          <List>
            {cartList && cartList.length > 0 &&
              cartList.map((cart) => (
                <CartListItem key={cart.id} cart={cart} isConfirm={true} />
              ))}
          </List>
        </div>
        <div className={classes.orderBox}>
          <TextDetail
            label={'商品合計'}
            value={'￥' + subtotal.toLocaleString()}
          />
          <TextDetail label={'消費税'} value={'￥' + tax.toLocaleString()} />
          <TextDetail
            label={'送料'}
            value={'￥' + shippingFee.toLocaleString()}
          />
          <Divider />
          <div className='module-spacer--extra-extra-small' />
          <TextDetail
            label={'合計(税込)'}
            value={'￥' + total.toLocaleString()}
          />
          <PrimaryButton label={'注文する'} onClick={order} />
        </div>
      </div>
    </section>
  );
};

const useStyles = makeStyles((theme) => ({
  // セクション1
  topSection: {
    padding: '40px 20px 0',
    [theme.breakpoints.up('md')]: {
      // PC
      padding: '60px 40px',
    },
  },
  topSectionTitle: {
    textAlign: 'center',
    fontWeight: 'bold',
    marginBottom: '32px',
    [theme.breakpoints.down('xs')]: {
      // SP
      fontSize: '20px',
    },
    [theme.breakpoints.up('sm')]: {
      // タブレット
      fontSize: '24px',
    },
    [theme.breakpoints.up('md')]: {
      // PC
      fontSize: '28px',
    },
  },
  detailBox: {
    margin: '0 auto',
    [theme.breakpoints.down('sm')]: {
      width: 320,
    },
    [theme.breakpoints.up('sm')]: {
      width: 512,
    },
  },
  orderBox: {
    border: '1px solid rgba(0, 0, 0, 0.2)',
    borderRadius: 4,
    boxShadow: '0 4px 2px 2px rgba(0, 0, 0, 0.2)',
    height: 256,
    margin: '24px auto 16px auto',
    padding: 16,
    width: 288,
  },
}));

export default OrderConfirm;
