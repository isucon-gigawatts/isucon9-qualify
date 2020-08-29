/*
getNewItems() 内の items 検索において created で検索される
select count(1) from items;
+----------+
| count(1) |
+----------+
|    50017 |
+----------+
1 row in set (0.09 sec)
*/
CREATE INDEX `hgsgtk_items_1` ON items (created_at);
/**
seller_id, created_at の検索 in getUserItems
*/
CREATE INDEX `hgsgtk_items_2` ON items (seller_id, created_at);
/**
buyer_id, created_at の検索 in getTransactions
*/
CREATE INDEX `hgsgtk_items_3` ON items (buyer_id, created_at);
/**
category_id, created_at の検索 in getNewCategoryItems
*/
CREATE INDEX `hgsgtk_items_4` ON items (category_id, created_at);

/**
ここまで
mysql> select TABLE_SCHEMA, TABLE_NAME, COLUMN_NAME, INDEX_NAME from INFORMATION_SCHEMA.STATISTICS where TABLE_SCHEMA='isucari';
+--------------+-----------------------+-------------------------+-----------------+
| TABLE_SCHEMA | TABLE_NAME            | COLUMN_NAME             | INDEX_NAME      |
+--------------+-----------------------+-------------------------+-----------------+
| isucari      | categories            | id                      | PRIMARY         |
| isucari      | configs               | name                    | PRIMARY         |
| isucari      | items                 | id                      | PRIMARY         |
| isucari      | items                 | category_id             | idx_category_id |
| isucari      | items                 | created_at              | hgsgtk_items_1  |
| isucari      | items                 | seller_id               | hgsgtk_items_2  |
| isucari      | items                 | created_at              | hgsgtk_items_2  |
| isucari      | items                 | buyer_id                | hgsgtk_items_3  |
| isucari      | items                 | created_at              | hgsgtk_items_3  |
| isucari      | items                 | category_id             | hgsgtk_items_4  |
| isucari      | items                 | created_at              | hgsgtk_items_4  |
| isucari      | shippings             | transaction_evidence_id | PRIMARY         |
| isucari      | transaction_evidences | id                      | PRIMARY         |
| isucari      | transaction_evidences | item_id                 | item_id         |
| isucari      | users                 | id                      | PRIMARY         |
| isucari      | users                 | account_name            | account_name    |
+--------------+-----------------------+-------------------------+-----------------+
16 rows in set (0.00 sec)

まだあまり効果はない
*/
