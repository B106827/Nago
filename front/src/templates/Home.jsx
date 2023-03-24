import { useEffect, useState } from 'react';
import { makeStyles } from '@material-ui/styles';
import { useDispatch, useSelector } from 'react-redux';
import { fetchProducts } from '../reducks/products/operations';
import { getProducts } from '../reducks/products/selectors';
import { ProductCard } from '../components/Products';

const Home = () => {
  const classes  = useStyles();
  const dispatch = useDispatch();
  const selector = useSelector((state) => state);

  useEffect(() => {
    dispatch(fetchProducts());
  }, []);

  const [products, setProducts] = useState([]);

  const fetchProductsList = getProducts(selector);
  useEffect(() => {
    if (fetchProductsList) setProducts(fetchProductsList);
  }, [fetchProductsList]);

  return (
    <section>
      <div className={classes.topSection}>
        <p className={classes.topSectionTitle}>商品一覧</p>
        <div className={classes.topSectionFlow}>
          {products.map((product, i) =>
            <div key={i}>
              <ProductCard
                id={product.id}
                name={product.name}
                images={product.images}
                price={product.price}
              />
            </div>
          )}
        </div>
      </div>
    </section>
  );
};

// スタイル
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
  topSectionFlow: {
    display: 'flex',
    flexWrap: 'wrap',
    flexDirection: 'row',
    '& > div': {
      padding: '16px 8px',
      [theme.breakpoints.down('xs')]: {
        // SP
        width: '100%',
      },
      [theme.breakpoints.up('sm')]: {
        // タブレット
        width: 'calc(100% / 2)',
      },
      [theme.breakpoints.up('md')]: {
        // PC
        width: 'calc(100% / 3)',
      },
    },
  },
}));

export default Home;
