import { fetchWrapper } from '../../utils/http';
import { showMessageAction } from '../messages/actions';
import { push } from 'connected-react-router';
import {
  fetchProductsAction,
  fetchProductAction,
} from './actions';

export const fetchProducts = () => {
  return async (dispatch) => {
    fetchWrapper(
      {
        type: 'GET',
        url: '/product',
      },
      dispatch,
    )
      .then((json) => {
        if (!json) return;
        if (json.status === 200) {
          const products = json.result.products;
          if (Array.isArray(products)) {
            dispatch(
              fetchProductsAction(products)
            );
          }
        } else {
          dispatch(showMessageAction('error', json.messages));
        }
      });
  };
};

export const fetchProduct = (productId) => {
  return async (dispatch) => {
    fetchWrapper(
      {
        type: 'GET',
        url: '/product/' + productId,
      },
      dispatch,
    )
      .then((json) => {
        if (!json) return;
        if (json.status === 200) {
          const product = json.result.product;
          if (product) {
            dispatch(
              fetchProductAction(product)
            );
          }
        } else {
          dispatch(showMessageAction('error', json.messages));
          dispatch(push('/'));
        }
      });
  };
}
