import { useDispatch } from 'react-redux';
import { push } from 'connected-react-router';
import Divider from '@material-ui/core/Divider';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import { makeStyles } from '@material-ui/styles';
import config from '../../config/base';

const OrderDetailListItem = (props) => {
  const classes  = useStyles();
  const dispatch = useDispatch();

  const detail               = props.detail;
  const num                  = detail.num;
  const price                = detail.price.toLocaleString();
  const priceWithTax         = detail.price * (1 + config.taxRate);
  const priceWithTax4Display = (detail.price * (1 + config.taxRate)).toLocaleString();
  const total                = (priceWithTax * num).toLocaleString();

  const product = detail.product;
  const image   = product.images[0]?.url;
  const name    = product.name;

  return (
    <>
      <ListItem className={classes.list}>
        <ListItemAvatar className='cursor-pointer' onClick={() => dispatch(push('/product/' + product.id))}>
          <img className={classes.image} src={image} alt='商品画像' />
        </ListItemAvatar>
        <div className={classes.text}>
          <ListItemText primary={name}/>
          <ListItemText primary={'￥' + price} secondary={<span className={classes.secondaryText}>{'個数:' + num}</span>}/>
          <ListItemText primary={'小計：' + priceWithTax4Display + '(税込)  × ' + num + ' = ' + total}/>
        </div>
      </ListItem>
      <Divider />
    </>
  );
};

const useStyles = makeStyles(() => ({
  list: {
    height: 128,
  },
  image: {
    objectFit: 'cover',
    margin: 16,
    height: 96,
    width: 96,
  },
  text: {
    width: '100%',
  },
  secondaryText: {
    display: 'block',
    textAlign: 'right',
    fontSize: '16px',
  }
}));


export default OrderDetailListItem;
