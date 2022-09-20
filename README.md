# img2txt
Convert  GIF, JP(E)G, or PNG image to text. Not very useful, but pretty cool.

# quick start

``` go
// go get github.com/frankxjkuang/img2txt

import "github.com/frankxjkuang/img2txt/src/img2txt"

i2t := img2txt.NewImg2txt()
i2t.SetPath("./look.png").PrintTxt()
```

![look](./src/img2txt/look.png)

```
                          #&::::o**&o:         .oooo::...        ....:oooo:::.             .o*&&:                     
                         .8*  oo.      :*8&:                                   .o&&*..o8&o.    *&.                    
                         .8o     .o&o                                                   .&8o   :8:                    
                         .8o         .o.                                              :8o      .8o                    
                         .8o                   .........                                       .8o                    
                       .&#8:                .............                                      .8:                    
                     .8&                  ....::::::......                                       *8:                  
                    *8.                  ....:::::::::........                                    .8&                 
                  .8o                   ....::::::::::::......                                      :8:               
                 :8:                   .....:::::::::::::.....                     ...:::..          .8o              
                o#.                     ...::::::::::::...........            ..::ooo:::..            .&*             
               o#.                     ..::ooooooooooooooo:::::....        ........                     &*            
              .8:                     ..:::ooo:::::::.....:::::::...        .......o888*o:.             .#o           
              o8                     ...:::::o::::. .   .....:::::..          .     &@&o:..              .8.          
              8&                    ...:::o*&&8#@#&o.   oWW8..::::..                                      8&          
              8*                    ...::::::o:::::::..       ........                                    *8.         
              8&                                               .....:..                                   :#o         
              &&                  ....  .....                   .::....                                   :#o         
              :8:               .................              :o:.....                                   :#o         
               *8.               ...................          o*o...:oo::..o*:                            *8.         
                &8               ....:::............           .:.....                                    8o          
                 &&.              ....::............            .....                                   .#o           
                  :#:               ...:.::::::::.......       ....:::o:.....                          :#:            
                    *#:               ...::::::::::.....     ..o**&&*o:..:::ooo:..                   o8o              
                      :#*               ....::::::::........:o&&&*oo:::::::...                     &8.                
                         :&&.               ....:::::::...:::::::....                          o8*.                   
                             :&&o.             .....:::::..........                      .o&&o.                       
                                   ::*&&*:.      .. ..........                 .:*&**:..                              
                                          :&o.:&8*::::o::::**o::....*8*:. .:&*.                                       
                                       .&&.       :*8o. o8:  .&& .&8:         :8*.                                                                                                                                                                                              
```