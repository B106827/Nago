import { useState } from 'react';
import Divider from '@material-ui/core/Divider';
import { TextDetail } from '../UIkit';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import Collapse from '@material-ui/core/Collapse';
import ExpandLess from '@material-ui/icons/ExpandLess';
import ExpandMore from '@material-ui/icons/ExpandMore';
import List from '@material-ui/core/List';
import { OrderDetailListItem } from './';

const datetimeToString = (date) => {
  return (
    date.getFullYear() +
    '-' +
    ('00' + (date.getMonth() + 1)).slice(-2) +
    '-' +
    ('00' + date.getDate()).slice(-2) +
    ' ' +
    ('00' + date.getHours()).slice(-2) +
    ':' +
    ('00' + date.getMinutes()).slice(-2) +
    ':' +
    ('00' + date.getSeconds()).slice(-2)
  );
};

const OrderHistoryItem = (props) => {
  const [orderedProductsOpen, setOrderedProductsOpen] = useState(false);

  const orderedProducsListClick = () => {
    setOrderedProductsOpen(!orderedProductsOpen);
  };

  const order     = props.order;
  const orderedAt = datetimeToString(new Date(order.orderedAt));
  const price     = '￥' + order.totalPrice.toLocaleString();

  return (
    <div>
      <div className='module-spacer--small' />
      <TextDetail label={'注文ID'} value={order.id} />
      <TextDetail label={'注文日時'} value={orderedAt} />
      <TextDetail label={'注文金額'} value={price} />
      <div>
        <ListItem button onClick={orderedProducsListClick}>
          <ListItemText primary="詳細を確認する" />
          {orderedProductsOpen ? <ExpandLess /> : <ExpandMore />}
        </ListItem>
        <Collapse in={orderedProductsOpen} timeout='auto' unmountOnExit>
          <List>
            {order.details && order.details.length > 0 && (
              order.details.map((detail) => (
                <OrderDetailListItem key={detail.id} detail={detail} />
              ))
            )}
          </List>
        </Collapse>
      </div>
      <div className='module-spacer--extra-extra-small' />
      <Divider />
    </div>
  );
};

export default OrderHistoryItem;
