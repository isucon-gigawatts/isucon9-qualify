/*
getNewItems() 内の items 検索において created で検索される
select count(1) from items;
+----------+
| count(1) |
+----------+
|    50017 |
+----------+
1 row in set (0.09 sec)

Order by 句があり、単独のcreatedでは、INDEXがきかないので、statusとの複合indexへ
参考: https://x1.inkenkun.com/archives/26

しかし、効果は今ひとつ 2010 -> 2010
*/
CREATE INDEX `hgsgtk_created_at` ON items (status, created_at);
