# comic-go

go rest api sample

mysql -u root -p -h host.docker.internal -P 29001 --protocol=tcp


LOAD DATA LOCAL INFILE "test.csv"
INTO TABLE comics
FIELDS TERMINATED BY ','
(@title,@number,@author,@imageUrl)
SET title=@title, volume=@number, author=@author, image_url=@imageUrl, created_at=CURRENT_TIMESTAMP, updated_at=CURRENT_TIMESTAMP;