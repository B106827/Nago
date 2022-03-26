//import { push } from 'connected-react-router';
//import { fetchWrapper } from '../utils/http';

//export const orderProduct = (productsInCart, amount) => {
//  return async (dispatch, getState) => {
//    const uid = getState().users.uid;
//    const userRef = db.collection('users').doc(uid);
//    const timestamp = FirebaseTimestamp.now();
//
//    let products = [],
//      soldOutProducts = [];
//
//    const batch = db.batch();
//
//    for (const product of productsInCart) {
//      const snapshot = await productsRef.doc(product.productId).get();
//      const sizes = snapshot.data().sizes;
//
//      const updatedSizes = sizes.map((size) => {
//        if (size.size === product.size) {
//          if (size.quantity === 0) {
//            soldOutProducts.push(product.name);
//            return size;
//          }
//          return {
//            size: size.size,
//            quantity: size.quantity - 1,
//          };
//        } else {
//          return size;
//        }
//      });
//
//      products.push({
//        id: product.productId,
//        images: product.images,
//        name: product.name,
//        price: product.price,
//        size: product.size,
//      });
//
//      batch.update(productsRef.doc(product.productId), { sizes: updatedSizes });
//
//      batch.delete(userRef.collection('cart').doc(product.cartId));
//    }
//
//    if (soldOutProducts.length > 0) {
//      const errorMessage =
//        soldOutProducts.length > 1
//          ? soldOutProducts.join('と')
//          : soldOutProducts[0];
//      alert(
//        '大変申し訳ありません。' +
//          errorMessage +
//          'が在庫切れとなったため、注文処理が中断しました'
//      );
//      return false;
//    } else {
//      batch
//        .commit()
//        .then(() => {
//          const orderRef = userRef.collection('orders').doc();
//          const date = timestamp.toDate();
//          // 配送日は3日後と設定
//          const shippingDate = FirebaseTimestamp.fromDate(
//            new Date(date.setDate(date.getDate() + 3))
//          );
//          const history = {
//            amount: amount,
//            created_at: timestamp,
//            id: orderRef.id,
//            products: products,
//            shipping_date: shippingDate,
//            updated_at: timestamp,
//          };
//          orderRef.set(history);
//          dispatch(push('/order/complete'));
//        })
//        .catch(() => {
//          alert(
//            '注文処理に失敗しました。通信環境をご確認の上、もう一度お試しください。'
//          );
//        });
//    }
//  };
//};
