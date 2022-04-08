# ES-Foxes

[![Latest Version](https://img.shields.io/github/v/release/kiwfy/es-foxes.svg?style=flat-square)](https://github.com/kiwfy/es-foxes/releases)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square&label=PRs%20Welcome)](http://makeapullrequest.com)



                                                       -:-` 
                                                      +mNNdo. 
                                             `       -NNNmNNm+`   
                                           `yhhys+:.`yNNmshNNNy` 
                                           /NNNNNNNmhNNNhssyNNNh.  
                                           :NNNyyhdmNNNNmmmNNNNNdyo/.``           
                                           `NNNhssssyhmNNNmddddmmNNNNdhyssssyyo-  
                                            yNNmsssssssyyssssssssyhdmmNNNNNNNNNN: 
                                            `mNNdssssssssssssssssssssssyhhdNNNNNy 
                                             -mNNmssssssssssssssyyhddmmNNNNNNNNNs 
                                              hNNmssssssssyhdmNNNNNNNNdhs++odNNm. 
                                             -NNNyssssydmNNNNNdhs+/-.``:ohNNNNs.  
                                             +NNNsssymNNNds+-``    `/ymNNNds/`    
                                             +NNNssdNNNs-`       .omNNNm+-`       
                                             :NNNymNNd-        `omNNmNNm/         
                                              dNNmNNm-        `yNNmhshNNNy-       
                                              -mNNNNd         oNNNysssymNNmy:`    
                                               /mNNNd        `mNNhssssssydNNNd+.  
                                                :mNNN:       `NNNyssssssssydmNNms:` 
                                                 -dNNd-...`` `dNNdssssssssssshmNNNh:` 
                                             `.:osdNNNmmmdddysdNNNhssssssssssssydNNNh- 
                                          `-ohmNNNmddNNNmmmmNNNNNNNdsssssssssssssydNNm+ 
                                        `/hNNNds/-.``hNNmssssyhdmNNNmhssssssssssssshNNNo  
                                       /dNNms-       -NNNysssssssydNNNmhssssssssssssyNNN/ 
                                     .yNNmo`          hNNdsssssssssshNNNNhsssssssssssdNNm 
                                    -mNNh.            yNNmsssssssssssshNNNmhsssssssssyNNN.
                                   :NNNs              dNNdssssssssssssssdNNNhsssssssssNNN:
                                  .mNNs              :NNNysssssssssssssssyhhsssssssssyNNN`
                                 `dNNy              `dNNdssssssssssssssssssssssssssssmNNy 
                                 oNNm.             `hNNmssssssssssssssssssssssssssssdNNm. 
                                `mNN+ `.-:/osyyhdddmNNNmdhyyssssssssssssssssssssssymNNd-  
                                +NNNsydmmNNNNmmdddddddmNNNNNmmdhyysssssssssssssyhmNNNs. 
                                hNNNNNdhs+/--..```````..-/oyhmNNNNmmddhhhhhhdmmNNNds-   
                                sddy/-`                     `.-/shdmmNNNNNNNNmdho:.       
                                 .`                               `.--:://::-.`      

                                    ▄▄▄▄▄▄▄▄         
                                    ██▀▀▀▀▀▀                                         
                                    ██         ▄████▄   ▀██  ██▀   ▄████▄   ▄▄█████▄ 
                                    ███████   ██▀  ▀██    ████    ██▄▄▄▄██  ██▄▄▄▄ ▀ 
                                    ██        ██    ██    ▄██▄    ██▀▀▀▀▀▀   ▀▀▀▀██▄ 
                                    ██        ▀██▄▄██▀   ▄█▀▀█▄   ▀██▄▄▄▄█  █▄▄▄▄▄██ 
                                    ▀▀          ▀▀▀▀    ▀▀▀  ▀▀▀    ▀▀▀▀▀    ▀▀▀▀▀▀  
                                                  

                                                  

This is a Golang script to backup, restore and transfer data between ElasticSearch clusters.

### Installation

Requires [Golang](https://go.dev) 1.12.1 or later 

Please create ```.env``` file from ```.env.example``` and fill in all keys, do the same for ```foxes.yaml``` file which should be created similar to ```foxes.yaml.example```


### Sample

All backed up data are stored in the backup folder in TXT file format, the first line contains the name of the ElasticSearch index and the other lines following the data of that index. Each line of the backup is in json format containing the event id and its source.

First it is necessary to build the project

```sh
go build ./ 
```

After that it is possible to execute the commands:

> **Backup**

```sh
./es-foxes backup index_name [url, default ELASTIC_URL]
```

examples:

```sh
./es-foxes backup event-2021-01-01
./es-foxes backup event-2021-01-01 http://127.0.0.1:9200
```


> **Restore**

```sh
./es-foxes restore file_name [index_to, default first line of file] [url, default ELASTIC_URL]
```

example:

```sh
./es-foxes restore event-2021-01-01.txt event-2021-01-02 http://127.0.0.1:9200 
```

>  **Copy**

```sh
./es-foxes restore url_from index_from url_to [index_to, default first line of file]
```

example:

```sh
./es-foxes copy http://127.0.0.1:9200 event-2021-01-01 http://127.0.0.1:9400 event-2021-01-02
```

> **Clear** 
    
You must have the **```foxes.yaml```** file configured

example:

```sh
./es-foxes clear
```

### Development

Want to contribute? Great!

The project using a simple code.
Make a change in your file and be careful with your updates!
**Any new code will only be accepted with all viladations.**


**Kiwfy - Open your code, open your mind!**