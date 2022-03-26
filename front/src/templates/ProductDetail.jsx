import { useState, useCallback } from 'react';
import { useDispatch } from 'react-redux';
import { makeStyles } from '@material-ui/styles';
import HTMLReactParser from 'html-react-parser';
import { ImageSwiper } from '../components/Products';
import { addProductToCart } from '../reducks/users/operations';
import { constants } from '../utils/constants';

const useStyles = makeStyles((theme) => ({
  sectionWrapperAdjuster: {
    marginTop: '50px',
  },
  sliderBox: {
    margin: '0 auto 24px auto',
    [theme.breakpoints.down('sm')]: {
      width: 320,
    },
    [theme.breakpoints.up('sm')]: {
      width: 400,
    },
  },
  productTitle: {
    fontWeight: 'bold',
  },
  detail: {
    textAlign: 'left',
    [theme.breakpoints.down('sm')]: {
      margin: '0 auto 16px auto',
      height: 'auto',
      width: 320,
    },
    [theme.breakpoints.up('sm')]: {
      margin: '0 auto',
      height: 'auto',
      width: 400,
    },
  },
  price: {
    fontSize: 36,
  },
}));

const returnCodeToBr = (text) => {
  if (text === '') {
    return text;
  } else {
    return HTMLReactParser(text.replace(/\r?\n/g, '<br/>'));
  }
};

const ProductDetail = () => {
  const classes = useStyles();
  const dispatch = useDispatch();
  //const selector = useSelector((state) => state);
  //const path = selector.router.location.pathname;

  const addProduct = useCallback();
  //(selectedSize) => {
  //  //const timestamp = FirebaseTimestamp.now();
  //  dispatch(
  //    addProductToCart({
  //      //added_at: timestamp,
  //      description: product.description,
  //      gender: product.gender,
  //      images: product.images,
  //      name: product.name,
  //      price: product.price,
  //      productId: product.id,
  //      quantity: 1,
  //    })
  //  );
  //},
  //[dispatch, product]
  const product = constants.product;

  return (
    <section
      className={'c-section-wrapin' + ' ' + `${classes.sectionWrapperAdjuster}`}
    >
      {product && (
        <div className='p-grid__row'>
          <div className={classes.sliderBox}>
            <ImageSwiper images={product.images} />
          </div>
          <div className={classes.detail}>
            <p>{product.subTitle}</p>
            <h2
              className={'u-text__headline' + ' ' + `${classes.productTitle}`}
            >
              {product.name}
            </h2>
            <p className={classes.price}>¥{product.price.toLocaleString()}</p>
            <div className='module-spacer--small' />
            <div className='module-spacer--small' />
            <p>{returnCodeToBr(product.description)}</p>
          </div>
        </div>
      )}
    </section>
  );
};

export default ProductDetail;
