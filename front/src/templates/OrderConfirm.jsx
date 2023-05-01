import { useState, useCallback, useMemo, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { makeStyles } from '@material-ui/styles';
import { CartListItem } from '../components/Products';
import { RegisterAddress } from '../components/Orders';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import Collapse from '@material-ui/core/Collapse';
import ExpandLess from '@material-ui/icons/ExpandLess';
import ExpandMore from '@material-ui/icons/ExpandMore';
import List from '@material-ui/core/List';
import Divider from '@material-ui/core/Divider';
import { PrimaryButton, TextDetail } from '../components/UIkit';
import { push } from 'connected-react-router';
import config from '../config/base';
import { getMyCartList } from '../reducks/users/selectors';
import { createOrder } from '../reducks/users/operations';
import { customValidErrResetAction } from '../reducks/utils/actions';
import { showMessageAction } from '../reducks/messages/actions';

const OrderConfirm = () => {
  const classes  = useStyles();
  const dispatch = useDispatch();
  const selector = useSelector((state) => state);
  const [cartListOpen, setCartListOpen] = useState(false);

  // カート
  const cartListClick = () => {
    setCartListOpen(!cartListOpen);
  };
  const cartList = getMyCartList(selector);
  useEffect(() => {
    if (!cartList || cartList.length === 0) {
      dispatch(push('/'));
    }
  })

  // お届け先
  const [address, setAddress] = useState({
    name: '',
    postcode: '',
    prefId: 0,
    primaryAddress: '',
    secondaryAddress: '',
    phoneNumber: '',
  });

  // 会計
  const subtotal = useMemo(() => {
    return cartList.reduce((sum, cart) => {
      return sum += (cart.product.price * cart.num);
    }, 0);
  }, [cartList]);
  // NOTE: 送料は仮で0円
  const shippingFee = 0;
  const tax         = subtotal * config.taxRate;
  const total       = subtotal + shippingFee + tax;
  const goToOrder = useCallback(() => {
    if (
      !address.name
      || !address.postcode
      || !address.prefId
      || !address.primaryAddress
      || !address.phoneNumber
    ) {
      dispatch(showMessageAction('error', 'お届け先に未入力項目があります'));
      return;
    }
    dispatch(customValidErrResetAction());
    dispatch(createOrder(total, address));
  }, [dispatch, address, total]);

  return (
    <section className={classes.topSection}>
      <h2 className={classes.topSectionTitle}>注文の確認</h2>
      <div className={'p-grid__row' + ' ' + `${classes.topSectionWrapper}`}>

        <div className={classes.cartBox}>
          {/* カートの中身（アコーディオン) */}
          <ListItem button onClick={cartListClick}>
            <ListItemText primary="カートの中身を確認する" />
            {cartListOpen ? <ExpandLess /> : <ExpandMore />}
          </ListItem>
          <Collapse in={cartListOpen} timeout='auto' unmountOnExit>
            <List>
              {cartList && cartList.length > 0 &&
                cartList.map((cart) => (
                  <CartListItem key={cart.id} cart={cart} isConfirm={true} />
                ))}
            </List>
          </Collapse>
          <div className='module-spacer--small' />

          {/* 届け先 */}
          <RegisterAddress address={address} setAddress={setAddress} />
        </div>

        {/* 会計 */}
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
          <PrimaryButton label={'支払い情報入力へ進む'} onClick={goToOrder} />
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
  topSectionWrapper: {
    justifyContent: 'center',
    maxWidth: theme.size.window.contentMaxWidth,
    margin: '0 auto',
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
  cartBox: {
    margin: '0 auto',
    [theme.breakpoints.down('sm')]: {
      width: 320,
    },
    [theme.breakpoints.up('sm')]: {
      width: 600,
    },
  },
  orderBox: {
    border: '1px solid rgba(0, 0, 0, 0.2)',
    borderRadius: 4,
    boxShadow: '0 4px 2px 2px rgba(0, 0, 0, 0.2)',
    height: 256,
    margin: '0 auto 16px auto',
    padding: 16,
    width: 300,
    [theme.breakpoints.down('sm')]: {
      marginTop: 16,
    },
  },
}));

export default OrderConfirm;
