import Divider from '@material-ui/core/Divider';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import { makeStyles } from '@material-ui/styles';
import DeleteIcon from '@material-ui/icons/Delete';
import IconButton from '@material-ui/core/IconButton';

const CartListItem = (props) => {
  const classes = useStyles();
  const cartId  = props.cart.id;
  const cartNum = props.cart.num;
  const product = props.cart.product;

  const image = product.images[0]?.url;
  const name  = product.name;
  const price = product.price.toLocaleString();

  const removeProductFromCart = () => {
    //         .collection('cart').doc(id)
    //         .delete();
  };

  return (
    <>
      <ListItem className={classes.list}>
        <ListItemAvatar>
          <img className={classes.image} src={image} alt='商品画像' />
        </ListItemAvatar>
        <div className={classes.text}>
          <ListItemText primary={name} secondary={<span className={classes.secondaryText}>{'個数:' + cartNum}</span>}/>
          <ListItemText primary={'￥' + price} />
        </div>
        <IconButton onClick={() => removeProductFromCart(cartId)}>
          <DeleteIcon />
        </IconButton>
      </ListItem>
      <Divider />
    </>
  );
};

const useStyles = makeStyles((theme) => ({
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
    [theme.breakpoints.down('xs')]: {
      // SP
      fontSize: '16px',
    },
    [theme.breakpoints.up('sm')]: {
      // SP
      fontSize: '18px',
    },
  }
}));


export default CartListItem;
