import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Typography from '@material-ui/core/Typography';
import NoImage from '../../assets/img/src/no_image.png';
import { push } from 'connected-react-router';
import { useDispatch } from 'react-redux';

const ProductCard = (props) => {
  const classes = useStyles();
  const dispatch = useDispatch();
  const images = props.images.length > 0 ? props.images : [{ path: NoImage }];
  const price = props.price.toLocaleString();

  return (
    <Card className={classes.card}>
      <CardMedia
        className={classes.media}
        image={images[0].path}
        onClick={() => dispatch(push('/product/' + props.id))}
      />
      <CardContent className={classes.content}>
        <div onClick={() => dispatch(push('/product/' + props.id))}>
          <Typography color='textSecondary' component='p'>
            {props.name}
          </Typography>
          <Typography align='right' component='p'>
            ￥{price}
          </Typography>
        </div>
      </CardContent>
    </Card>
  );
};

// スタイル
const useStyles = makeStyles(() => ({
  card: {
    padding: 0,
  },
  // 商品画像
  media: {
    paddingTop: '100%',
  },
  content: {
    padding: '12px',
    paddingBottom: '12px !important',
    textAlign: 'left',
    '& > div': {
      overflow: 'hidden',
      textOverflow: 'ellipsis',
      whiteSpace: 'nowrap',
    },
  },
}));


export default ProductCard;
