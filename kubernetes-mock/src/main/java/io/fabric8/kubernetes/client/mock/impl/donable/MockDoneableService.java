/**
 * Copyright (C) 2015 Red Hat, Inc.                                        
 *                                                                         
 * Licensed under the Apache License, Version 2.0 (the "License");         
 * you may not use this file except in compliance with the License.        
 * You may obtain a copy of the License at                                 
 *                                                                         
 *         http://www.apache.org/licenses/LICENSE-2.0                      
 *                                                                         
 * Unless required by applicable law or agreed to in writing, software     
 * distributed under the License is distributed on an "AS IS" BASIS,       
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and     
 * limitations under the License.
 */

package io.fabric8.kubernetes.client.mock.impl.donable;

import io.fabric8.kubernetes.api.builder.Visitor;
import io.fabric8.kubernetes.api.model.Doneable;
import io.fabric8.kubernetes.api.model.DoneableService;
import io.fabric8.kubernetes.api.model.Service;
import io.fabric8.kubernetes.api.model.ServiceFluent;
import io.fabric8.kubernetes.api.model.Service;
import io.fabric8.kubernetes.api.model.ServiceFluent;
import io.fabric8.kubernetes.api.model.ServiceFluentImpl;
import io.fabric8.kubernetes.client.mock.MockDoneable;
import org.easymock.EasyMock;
import org.easymock.IExpectationSetters;

public class MockDoneableService extends ServiceFluentImpl<MockDoneableService> implements MockDoneable<Service> {

  private interface DelegateInterface extends Doneable<Service>, ServiceFluent<DoneableService> {}
  private final Visitor<Service> visitor = new Visitor<Service>() {
    @Override
    public void visit(Service item) {}
  };

  private final DelegateInterface delegate;

  public MockDoneableService() {
    this.delegate = EasyMock.createMock(DelegateInterface .class);
  }

  @Override
  public IExpectationSetters<Service> done() {
    return EasyMock.expect(delegate.done());
  }

  @Override
  public Void replay() {
    EasyMock.replay(delegate);
    return null;
  }

  @Override
  public void verify() {
    EasyMock.verify(delegate);
  }

  @Override
  public Doneable<Service> getDelegate() {
    return new DoneableService(visitor) {
      @Override
      public Service done() {
        return delegate.done();
      }
    };
  }
}
